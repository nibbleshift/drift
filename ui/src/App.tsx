import { useState, useEffect } from 'react';
import {
  MantineProvider,
  AppShell,
  Navbar,
  Header,
  Footer,
  Aside,
  Text,
  TextInput,
  Group,
  Badge,
  Card,
  MediaQuery,
  Burger,
  List,
  Stack,
  useMantineTheme,
  createStyles,
  rem,
  Code,
  ScrollArea,
} from '@mantine/core';

import { Route, BrowserRouter, Routes } from "react-router-dom";

import SDK from 'casdoor-js-sdk';

import {
  IconNotes,
  IconCalendarStats,
  IconGauge,
  IconPresentationAnalytics,
  IconFileAnalytics,
  IconAdjustments,
  IconLock,
} from '@tabler/icons-react';
import { LinksGroup } from './NavbarLinksGroup';
import { Logo } from './Logo';
import { UserButton } from './UserButton';
import { Home } from './Home';
import { useQuery, useMutation, gql } from '@apollo/client';

import { config } from "./Setting";
import { AuthCallback } from "casdoor-react-sdk";

const GET_USERS = gql`
query GetUsers {
  users {
    edges {
      node {
        id
        username
        firstName
        lastName
        posts {
	  createdAt
          data
        }
      }
    }
  }
}
`;

const useStyles = createStyles((theme) => ({
  navbar: {
    backgroundColor: theme.colorScheme === 'dark' ? theme.colors.dark[6] : theme.white,
    paddingBottom: 0,
  },

  header: {
    padding: theme.spacing.md,
    paddingTop: 0,
    marginLeft: `calc(${theme.spacing.md} * -1)`,
    marginRight: `calc(${theme.spacing.md} * -1)`,
    color: theme.colorScheme === 'dark' ? theme.white : theme.black,
    borderBottom: `${rem(1)} solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
    }`,
  },

  links: {
    marginLeft: `calc(${theme.spacing.md} * -1)`,
    marginRight: `calc(${theme.spacing.md} * -1)`,
  },

  linksInner: {
    paddingTop: theme.spacing.xl,
    paddingBottom: theme.spacing.xl,
  },

  footer: {
    marginLeft: `calc(${theme.spacing.md} * -1)`,
    marginRight: `calc(${theme.spacing.md} * -1)`,
    borderTop: `${rem(1)} solid ${
      theme.colorScheme === 'dark' ? theme.colors.dark[4] : theme.colors.gray[3]
    }`,
  },
}));

const mockdata = [
  { label: 'Dashboard', icon: IconGauge },
  {
    label: 'Market news',
    icon: IconNotes,
    initiallyOpened: true,
    links: [
      { label: 'Overview', link: '/' },
      { label: 'Forecasts', link: '/' },
      { label: 'Outlook', link: '/' },
      { label: 'Real time', link: '/' },
    ],
  },
  {
    label: 'Releases',
    icon: IconCalendarStats,
    links: [
      { label: 'Upcoming releases', link: '/' },
      { label: 'Previous releases', link: '/' },
      { label: 'Releases schedule', link: '/' },
    ],
  },
  { label: 'Analytics', icon: IconPresentationAnalytics },
  { label: 'Contracts', icon: IconFileAnalytics },
  { label: 'Settings', icon: IconAdjustments },
  {
    label: 'Security',
    icon: IconLock,
    links: [
      { label: 'Enable 2FA', link: '/' },
      { label: 'Change password', link: '/' },
      { label: 'Recovery codes', link: '/' },
    ],
  },
];


function LeftNav() {
  const { classes } = useStyles();
  const links = mockdata.map((item) => <LinksGroup {...item} key={item.label} />);

  return (
    <Navbar height={800} width={{ sm: 300 }} p="md" className={classes.navbar}>
      <Navbar.Section grow className={classes.links} component={ScrollArea}>
        <div className={classes.linksInner}>{links}</div>
      </Navbar.Section>

      <Navbar.Section className={classes.footer}>
        <UserButton
          image="https://images.unsplash.com/photo-1508214751196-bcfd4ca60f91?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=255&q=80"
          name="Ann Nullpointer"
          email="anullpointer@yahoo.com"
        />
      </Navbar.Section>
    </Navbar>
  );
}


export default function App() {
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);
  const [username, setUsername] = useState('');
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [tokenReceived, setTokenReceived] = useState(false);
  const [sdk, setSdk] = useState(new SDK(config));

  useEffect(() => {
    if (window.location.href.indexOf('code') !== -1) {
      if (!sessionStorage.getItem('token')) {
        sdk.signin("http://localhost:8000").then(res => {
          sessionStorage.setItem('token', res.token);
          setTokenReceived(true);
        });
      }
    }
  }, [sdk]);

  useEffect(() => {
    if (sessionStorage.getItem('token')) {
      getInfo().then(res => setInfo(res));

      async function getInfo() {
        let token = sessionStorage.getItem('token');
        if (!token) {
          return;
        }
        else {
          return fetch(`http://localhost:9900/api/getUserInfo?token=${token}`).then(res => res.json());
        }
      }

      function setInfo(res) {
        let userinfo = res;
        setUsername(userinfo.name);
        setIsLoggedIn(true);
      }
    }
  }, [tokenReceived])

  const authCallback = (
    <AuthCallback
      sdk={SDK}
      serverUrl={config.ServerUrl}
      saveTokenFromResponse={(res) => {
        // @ts-ignore
        // save token
        localStorage.setItem("token", res.data.accessToken);
      }}
      isGetTokenSuccessful={(res) => {
        // @ts-ignore
        // according to the data returned by the server,
        // determine whether the `token` is successfully obtained through `code` and `state`.
        return res.success === true;
      }}
    />
  );

  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
    {
      isLoggedIn
      ? <AppShell
      styles={{
        main: {
          background: theme.colorScheme === 'dark' ? theme.colors.dark[8] : theme.colors.gray[0],
        },
      }}
      navbarOffsetBreakpoint="sm"
      asideOffsetBreakpoint="sm"
      navbar={ 
        LeftNav()
      }
      aside={
        <MediaQuery smallerThan="sm" styles={{ display: 'none' }}>
          <Aside p="md" hiddenBreakpoint="sm" width={{ sm: 200, lg: 300 }}>
          </Aside>
        </MediaQuery>
      }
      header={
        <Header height={{ base: 50, md: 70 }} p="md">
          <div style={{ display: 'flex', alignItems: 'center', height: '100%' }}>
            <MediaQuery largerThan="sm" styles={{ display: 'none' }}>
              <Burger
                opened={opened}
                onClick={() => setOpened((o) => !o)}
                size="sm"
                color={theme.colors.gray[6]}
                mr="xl"
              />
            </MediaQuery>
            <Text>DriftSocial</Text>
          </div>
        </Header>
      }
    >
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home/>} />
      </Routes>
    </BrowserRouter>
    </AppShell> : window.location.href = sdk.getSigninUrl()
    }
    </MantineProvider> 
  );
}
