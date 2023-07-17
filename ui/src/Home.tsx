import { Drifts } from './Drifts';
import { config } from './Setting';
import { SilentSignin, isSilentSigninRequired } from "casdoor-react-sdk";
import SDK from 'casdoor-js-sdk';
import React, { useState, useEffect } from 'react';

export function Home() {
  const isLoggedIn = () => {
    return localStorage.getItem("token") !== null;
  };

const [sdk, setSdk] = useState(new SDK(config));

  if (isSilentSigninRequired()) {
    // if the `silentSignin` parameter exists, perform silent login processing
    return (
      <SilentSignin
        sdk={Setting.CasdoorSDK}
        isLoggedIn={isLoggedIn}
        handleReceivedSilentSigninSuccessEvent={() => {
          // jump to the home page here and clear silentSignin parameter
          window.location.href = "/";
        }}
        handleReceivedSilentSigninFailureEvent={() => {
          // prompt the user to log in failed here
          alert("login failed");
        }}
      />
    );
  }

  const renderContent = () => {
    if (isLoggedIn()) {
      return <Drifts />;
    } else {
      return <div>not logged in</div>
    }
  };

  return renderContent();
}

