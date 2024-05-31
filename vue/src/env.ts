const apiURL = import.meta.env.VITE_API_COSMOS ?? "https://api.warmstorage.xyz";
const rpcURL = import.meta.env.VITE_WS_TENDERMINT ?? "https://rpc.warmstorage.xyz";
const prefix = import.meta.env.VITE_ADDRESS_PREFIX ?? "cosmos";

export const env = {
  apiURL,
  rpcURL,
  prefix,
};
