import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { AuthProvider } from "oidc-react";

import {
  ApolloProvider,
  ApolloClient,
  createHttpLink,
  InMemoryCache
} from '@apollo/client';


const httpLink = createHttpLink({
  uri: 'http://localhost:8081/query'
});

const client = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache()
});

const oidcConfig = {
  onSignIn: async (response: any) => {
    console.log(response.profile);
    window.location.hash = "";
  },
  authority: "http://localhost:8080", // replace with your instance
  //clientId: "223672648771633155@drift",
  clientId: "223673647603187715@drift",
  responseType: "code",
  redirectUri: "http://localhost:5173/callback",
  scope: "openid profile email",
};

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
   <AuthProvider {...oidcConfig}>
    <ApolloProvider client={client}>
      <App />
      </ApolloProvider>
   </AuthProvider>
  </React.StrictMode>
)
