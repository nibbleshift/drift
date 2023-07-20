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
    window.location.hash = "";
    window.location = "/";
    console.log(response.profile);
  },
  authority: "http://localhost:8080",
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
