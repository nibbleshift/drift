import { Route, BrowserRouter, Routes } from "react-router-dom";
import {Home} from "./Home";
import { AuthCallback } from "casdoor-react-sdk";
import * as Setting from "./Setting";
import { SilentSignin, isSilentSigninRequired } from "casdoor-react-sdk";

export default function AuthProvider() {
  const authCallback = (
    <AuthCallback
      sdk={Setting.CasdoorSDK}
      serverUrl={Setting.ServerUrl}
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
    <BrowserRouter>
      <Routes>
        <Route path="/callback" element={authCallback} />
        <Route path="/" element={<Home />} />
      </Routes>
    </BrowserRouter>
    );
  }
