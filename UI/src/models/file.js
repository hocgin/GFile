import API from '@/util/api';
import { message } from 'antd';
import * as qs from 'qs';

export default {
  namespace: 'file',
  state: {
    page: [],
  },
  effects: {
    * page({ payload }, { call, put, select }) {
      if (payload.page === 1) {
        yield put({
          type: 'fillPage',
          payload: [],
        });
      }

      let result = yield API.searchFile(payload);
      if (result.code === 200) {
        let page = yield select(({file}) => file.page);
        yield put({
          type: 'fillPage',
          payload: [...(page || []), ...(result.data || [])],
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
        all: payload,
      };
    },
    fillPage(state, { payload }) {
      return {
        ...state,
        page: payload,
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
              type: 'page',
              payload: {
                page: 1,
                size: 15,
              },
            });
            dispatch({
              type: 'tag/all',
              payload: {
                page: 1,
                size: 15,
              },
            });
            break;
          }
          default:
        }
      });
    },
  },
};