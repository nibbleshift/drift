import {
  Text,
  Card,
  Group,
  Badge,
  TextInput,
  Stack,
  Loader,
} from '@mantine/core';


import TimeAgo from 'javascript-time-ago'
import en from 'javascript-time-ago/locale/en.json'
TimeAgo.addDefaultLocale(en)
import ReactTimeAgo from 'react-time-ago'

import { useQuery, useMutation, gql } from '@apollo/client';

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

const CREATE_POST2 = gql`
mutation createPost2($id: Int!, $post: String!, $tags: [String], $mentions: [String]) {
  createPost2(id: $id, post: $post, tags: $tags, mentions: $mentions) {
    id
  }
}
`;

function DisplayPosts() {
  const { loading, error, data } = useQuery(GET_POSTS);

  if (loading) return <Loader />;
  if (error) return <p>Error : {error.message}</p>;

   var cards = data.posts.edges.map(({ node }) => (
    <Card shadow="sm" padding="xs" radius="xs" withBorder key={node.id}>
      <Group position="apart" mt="xs" mb="xs">
        <Group position="left" mt="xs" mb="xs">
          <Text weight={500} fz="sm">{node.owner.username}</Text>
          <Text weight={500} fz="sm" c="dimmed"><ReactTimeAgo date={Date.parse(node.createdAt)} locale="en-US"/></Text>
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
  var items = cards.map((item, idx) => (
    <div key={idx}>{item}</div>
  ));

  return <Stack spacing="xs">{items}</Stack>;
}

function processPost(msg) {
  const tags = msg.match(/\#[a-zA-Z0-9_]+/g);
  const mentions = msg.match(/\@[a-zA-Z0-9_]+/g);

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

  return { tags, mentions };
}

export function Drifts() {
  const [addPost2 ] = useMutation(CREATE_POST2);

    return (
    <div>
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
          const { tags, mentions } = processPost(msg);
          addPost2({ variables: {id: 1, post: msg, tags: tags, mentions: mentions} });
          e.currentTarget.value = '';
        }
      }}
    />
    <DisplayPosts />
    </div>
    );
}
