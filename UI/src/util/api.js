import request from '@/util/request';
import { stringify } from 'qs';

export default class API {

  static worked(payload) {
    return request(`/worked`, {
      method: 'GET',
      body: {
        ...payload,
      },
    });
  }

  /**
   * 查询所有文件列表
   * @param payload
   * @returns {*}
   */
  static queryAllFile(payload) {
    let query = stringify(payload);
    return request(`/file?${query}`, {
      method: 'GET',
    });
  }

  /**
   * 查询所有文件列表
   * @param payload
   * @returns {*}
   */
  static searchFile(payload) {
    let query = stringify(payload);
    return request(`/file/_search?${query}`, {
      method: 'GET',
    });
  }

  /**
   * 查询路径列表
   * @param payload
   * @returns {*}
   */
  static queryAllTag(payload) {
    let query = stringify(payload);
    return request(`/classify?${query}`, {
      method: 'GET',
    });
  }
}