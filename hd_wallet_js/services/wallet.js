import { Wallet, Mnemonic } from "ethers";

export default class EthWallet {
  constructor() {
    this.pathBIP = "m/44'/60'/0'/0/";
  }

  restoreFromPrivateKey(pk) {
    const wallet = new Wallet(pk);
    const keys = {
      address: wallet.address,
      privateKey: wallet.privateKey,
    };
    return { seedPhrase: null, keys };
  }

  restoreFromMnemonic(mnemonic, limit) {
    if (!Mnemonic.isValidMnemonic(mnemonic))
      throw new Error(`Invalid seed phrase: [${mnemonic}]`);
    const wallet = Wallet.fromPhrase(mnemonic);
    const keys = [...Array(limit).keys()].map((i) => {
      const hdNodeWallet = wallet.derivePath(this.pathBIP + i);
      return {
        address: hdNodeWallet.address,
        privateKey: hdNodeWallet.privateKey,
      };
    });
    return { seedPhrase: wallet.mnemonic.phrase, keys };
  }

  generate() {
    const {
      mnemonic: { phrase },
      address,
      privateKey,
    } = Wallet.createRandom();
    return { seedPhrase: phrase, keys: [{ address, privateKey }] };
  }
}
