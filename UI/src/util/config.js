export default class Config {
  static host() {
    if (Config.isDev()) {
      return 'http://127.0.0.1:8080';
    }
    return 'http://192.168.1.217:8080';
  }

  static isDev() {
    return false;
  }
}