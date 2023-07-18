import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

import { AuthCallback } from "casdoor-react-sdk";

import {
  ApolloProvider,
  ApolloClient,
  createHttpLink,
  InMemoryCache
} from '@apollo/client';


const httpLink = createHttpLink({
  uri: 'http://localhost:8080/query'
});

const client = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache()
});

function gotoSignup() {
  window.location.href = sdk.getSigninUrl();
}

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </React.StrictMode>,
)
