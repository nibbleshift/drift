import Sdk from "casdoor-js-sdk";

export const ServerUrl = "http://localhost:7023";

const sdkConfig = {
  serverUrl: "http://localhost:8000",
  clientId: "586d88470bc11dbce0b6",
  appName: "application_leo4ii",
  organizationName: "organization_1110ds",
  redirectPath: "/callback",
  signinPath: "/api/signin",
};

export const CasdoorSDK = new Sdk(sdkConfig);
