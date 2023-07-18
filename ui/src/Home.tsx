import { Drifts } from './Drifts';
import { config } from './Setting';
import { SilentSignin, isSilentSigninRequired } from "casdoor-react-sdk";
import SDK from 'casdoor-js-sdk';
import React, { useState, useEffect } from 'react';

export function Home() {
  const renderContent = () => {
      return <Drifts />;
  };

  return renderContent();
}

