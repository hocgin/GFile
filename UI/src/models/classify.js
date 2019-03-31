import * as qs from 'qs';
import API from '@/util/api';
import { message } from 'antd';

export default {
  namespace: 'classify',
  state: {
    all: [],
  },
  effects: {
    * queryAll({ payload }, { call, put }) {
      let result = yield API.queryAllTag(payload);
      if (result.code === 200) {
        yield put({
          type: 'fillAll',
          payload: result.data || [],
        });
      } else {
        message.error(result.message);
      }
    },
  },
  reducers: {
    fillAll(state, { payload }) {
      return {
        ...state,
        all: Object.keys(payload),
      };
    },
  },
  subscriptions: {
    setup({ dispatch, history }, done) {
      // 监听路由的变化，请求页面数据
      return history.listen(({ pathname, search }) => {
        const query = qs.parse(search);
        switch (pathname) {
          case '/': {
            dispatch({
              type: 'queryAll',
              payload: {},
            });
            break;
          }
          default:
        }
      });
    },
  },
};