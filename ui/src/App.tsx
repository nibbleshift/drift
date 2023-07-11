import { useState } from 'react';
import {
  MantineProvider,
  AppShell,
  Navbar,
  Header,
  Footer,
  Aside,
  Text,
  Group,
  Badge,
  Card,
  MediaQuery,
  Burger,
  List,
  Stack,
  useMantineTheme,
} from '@mantine/core';

import { useQuery, gql } from '@apollo/client';

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
  posts {
    edges {
      node {
        id
        data
        owner {
          username
        }
      }
    }
  }
}
`;

function DisplayPosts() {
  const { loading, error, data } = useQuery(GET_POSTS);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

   var cards = data.posts.edges.map(({ node }) => (
    <Card shadow="sm" padding="xs" radius="xs" withBorder key={node.id}>
      <Group position="apart" mt="xs" mb="xs">
        <Text weight={500}>{node.owner.username}</Text>
        <Badge color="green" variant="light">
          Verified
        </Badge>
      </Group>
      <Text size="sm" color="dimmed">
        {node.data}
      </Text>
    </Card>
    
  ));
  var items = cards.map((item) => (
    <div key={item}>{item}</div>
  ));

  return <Stack spacing="xs">{items}</Stack>;
}

export default function App() {
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);
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
        <Navbar p="md" hiddenBreakpoint="sm" hidden={!opened} width={{ sm: 200, lg: 300 }}>
          <Navbar.Section>Test</Navbar.Section>
        </Navbar>
      }
      aside={
        <MediaQuery smallerThan="sm" styles={{ display: 'none' }}>
          <Aside p="md" hiddenBreakpoint="sm" width={{ sm: 200, lg: 300 }}>
            <Text>Application sidebar</Text>
          </Aside>
        </MediaQuery>
      }
      footer={
        <Footer height={60} p="md">
          Application footer
        </Footer>
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
      <DisplayPosts />
    </AppShell>
    </MantineProvider>
  );
}
