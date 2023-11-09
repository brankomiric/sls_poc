import HDWallet from "./services/wallet.js";

const wallet = new HDWallet();

export const generate = (params) => {
  const limit = +params?.accountLimit || 10;
  const { seedPhrase } = wallet.generate();
  return wallet.restoreFromMnemonic(seedPhrase, limit);
};

export const restore = (body) => {
  const limit = body.accountLimit || 10;
  const seedPhrase = body.seedPhrase;
  return wallet.restoreFromMnemonic(seedPhrase, limit);
};
