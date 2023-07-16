import { useState } from 'react';
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


import emojify from 'emojify';

import TimeAgo from 'javascript-time-ago'
import en from 'javascript-time-ago/locale/en.json'
TimeAgo.addDefaultLocale(en)
import ReactTimeAgo from 'react-time-ago'

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
import { useQuery, useMutation, gql } from '@apollo/client';

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

const GET_POSTS = gql`
query GetPosts {
  posts(first: 10, orderBy: {direction: DESC, field: CREATED_AT}) {
    edges {
      node {
        id
        createdAt
        data
        owner {
          username
        }
      }
    }
  }
}
`;

const CREATE_POST = gql`
mutation createPost($data: String!) {
  createPost(input:{data: $data, ownerID:7}) {
    id
  }
}
`;

const CREATE_POST2 = gql`
mutation createPost2($id: Int!, $post: String!, $tags: [String], $mentions: [String]) {
  createPost2(id: $id, post: $post, tags: $tags, mentions: $mentions) {
    id
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

function DisplayPosts() {
  const { loading, error, data } = useQuery(GET_POSTS);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

   var cards = data.posts.edges.map(({ node }) => (
    <Card shadow="sm" padding="xs" radius="xs" withBorder key={node.id}>
      <Group position="apart" mt="xs" mb="xs">
        <Group position="left" mt="xs" mb="xs">
          <Text weight={500} fz="sm">{node.owner.username}</Text>
          <Text weight={500} fz="sm" c="dimmed"><ReactTimeAgo date={node.createdAt} locale="en-US"/></Text>
        </Group>
          <Badge color="green" variant="light">
            Verified
          </Badge>
        </Group>
      <Text size="sm" >
        {node.data}
      </Text>
    </Card>
    
  ));
  var items = cards.map((item) => (
    <div key={item}>{item}</div>
  ));

  return <Stack spacing="xs">{items}</Stack>;
}

function processPost(msg) {
  const tags = msg.match(/\#[a-zA-Z0-9_]+/g);
  const mentions = msg.match(/\@[a-zA-Z0-9_]+/g);
  const emojis = msg.match(/[\u{1f300}-\u{1f5ff}\u{1f900}-\u{1f9ff}\u{1f600}-\u{1f64f}\u{1f680}-\u{1f6ff}\u{2600}-\u{26ff}\u{2700}-\u{27bf}\u{1f1e6}-\u{1f1ff}\u{1f191}-\u{1f251}\u{1f004}\u{1f0cf}\u{1f170}-\u{1f171}\u{1f17e}-\u{1f17f}\u{1f18e}\u{3030}\u{2b50}\u{2b55}\u{2934}-\u{2935}\u{2b05}-\u{2b07}\u{2b1b}-\u{2b1c}\u{3297}\u{3299}\u{303d}\u{00a9}\u{00ae}\u{2122}\u{23f3}\u{24c2}\u{23e9}-\u{23ef}\u{25b6}\u{23f8}-\u{23fa}]/ug);

  if (tags != null) {
    for (var i = 0; i < tags.length; i++) {
      tags[i] = tags[i].slice(1);
    }
  }

  if (mentions != null) {
    for (var i = 0; i < mentions.length; i++) {
      mentions[i] = mentions[i].slice(1);
    }
  }

  if (emojis != null) {
    for (var i = 0; i < emojis.length; i++) {
      emojis[i] = emojis[i].slice(1);
    }
  }

  return { tags, mentions, emojis };
}

export default function App() {
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);
  //const [addPost, { data, loading, error }] = useMutation(CREATE_POST);
  const [addPost2, { data, loading, error }] = useMutation(CREATE_POST2);
  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
    <AppShell
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
    <TextInput
      radius="md"
      size="md"
      placeholder="What's on your mind?"
      onKeyDown={e => {
        if (e.key === 'Enter') {
          var msg = e.currentTarget.value.trim();

	  if (msg === "") {
            return;
          }

          e.preventDefault();
          const { tags, mentions, emojis } = processPost(msg);
          console.log(tags);
          console.log(mentions);
          console.log(emojis);
          //addPost({ variables: {data: msg} });
          addPost2({ variables: {id: 1, post: msg, tags: tags, mentions: mentions} });
          e.currentTarget.value = '';
        }
      }}
    />
      <DisplayPosts />
    </AppShell>
    </MantineProvider>
  );
}
