import styles from './index.less';
import React from 'react';
import { connect } from 'dva';
import { Divider, List, Select } from 'antd';
import InfiniteScroll from 'react-infinite-scroller';
import DPlayer from 'react-dplayer';

let page = 1;

@connect(({ file, classify, loading }) => {
  return {
    data: file.page,
    allClassify: ['所有', ...classify.all],
    isLoading: loading.effects['file/page'],
  };
}, (dispatch) => ({
  $page: (args = {}) => dispatch({ type: 'file/page', ...args }),
}))
class IndexLayout extends React.Component {

  state = {
    url: null,
  };

  player = null;

  componentDidMount() {

  }

  render() {
    let { data, allClassify, isLoading } = this.props;
    let { url } = this.state;

    return (
      <div className={styles.page}>
        {/*<h1 className={styles.title}/>*/}
        <div className={styles.videoBox}>
          <DPlayer style={{ height: '100%' }}
                   video={{
                     url: url,
                   }}
                   onLoad={this.onLoadPlayer}
          />
        </div>
        <div>
          <Divider orientation={'left'}>目录</Divider>
          <Select defaultValue="请选择类别" style={{ width: 120 }}
                  onChange={this.onClassifyChange}>
            {(allClassify || []).map((item) => {
              return (<Select.Option value={item}>{item}</Select.Option>);
            })}
          </Select>
        </div>
        <div className={styles.listContainer}>
          <InfiniteScroll initialLoad={false}
                          loadMore={this.onLoadMore}
                          hasMore={!isLoading}
                          useWindow={false}>
            <List itemLayout="horizontal"
                  size="large"
                  loading={isLoading}
                  dataSource={data}
                  renderItem={this.renderListItem}>
            </List>
          </InfiniteScroll>
        </div>
      </div>
    );
  }

  onLoadPlayer = (player) => {
    this.player = player;
  };

  // 渲染项
  renderListItem = (item) => {

    let actions = {
      'video': [(<a onClick={this.onClickPlayAction.bind(this, item)}>播放</a>), (
        <a onClick={this.onClickLoadingAction.bind(this, item)}>加载</a>)],
    };
    let defaultAction = [(<a onClick={this.onClickDefaultAction.bind(this, item)}>查看内容</a>)];
    let action = actions[item.typeOf] || [defaultAction];
    return (<List.Item actions={[...action]}>
      <List.Item.Meta title={item.fileName}
                      description={item.tags}/>
    </List.Item>);
  };

  // 加载更多
  onLoadMore = () => {
    let { $page } = this.props;
    $page({
      payload: {
        size: 10,
        page: ++page,
      },
    });
  };

  // 播放
  onClickPlayAction = (item) => {
    this.player.switchVideo({
      url: item.path,
    });
    this.player.play();
  };

  // 加载
  onClickLoadingAction = (item) => {
    this.player.switchVideo({
      url: item.path,
    });
  };

  // 查看内容
  onClickDefaultAction = (item) => {
    window.open(item.path);
  };

  // 改变分类
  onClassifyChange = (v) => {
    let { $page } = this.props;
    page = 1;
    $page({
      payload: {
        size: 10,
        page: page,
        classify: v
      },
    });
  };

}

export default IndexLayout;