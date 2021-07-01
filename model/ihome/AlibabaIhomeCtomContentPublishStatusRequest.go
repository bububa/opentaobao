package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
实拍图发布审核状态查询API API请求
alibaba.ihome.ctom.content.publish.status

实拍图发布审核状态查询API
*/
type AlibabaIhomeCtomContentPublishStatusAPIRequest struct {
    model.Params
    // 要查询投稿状态的ID列表
    _idList   []int64
}

// 初始化AlibabaIhomeCtomContentPublishStatusAPIRequest对象
func NewAlibabaIhomeCtomContentPublishStatusRequest() *AlibabaIhomeCtomContentPublishStatusAPIRequest{
    return &AlibabaIhomeCtomContentPublishStatusAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIhomeCtomContentPublishStatusAPIRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.content.publish.status"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIhomeCtomContentPublishStatusAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdList Setter
// 要查询投稿状态的ID列表
func (r *AlibabaIhomeCtomContentPublishStatusAPIRequest) SetIdList(_idList []int64) error {
    r._idList = _idList
    r.Set("id_list", _idList)
    return nil
}

// IdList Getter
func (r AlibabaIhomeCtomContentPublishStatusAPIRequest) GetIdList() []int64 {
    return r._idList
}
