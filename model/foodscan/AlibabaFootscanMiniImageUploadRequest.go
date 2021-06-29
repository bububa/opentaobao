package foodscan

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家端图片上传 API请求
alibaba.footscan.mini.image.upload

提供图片上传功能，同时进行图片的检测
*/
type AlibabaFootscanMiniImageUploadRequest struct {
    model.Params
    // 平台分配的token
    token   string
    // 请求数据
    reqData   *CheckParam
}

// 初始化AlibabaFootscanMiniImageUploadRequest对象
func NewAlibabaFootscanMiniImageUploadRequest() *AlibabaFootscanMiniImageUploadRequest{
    return &AlibabaFootscanMiniImageUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniImageUploadRequest) GetApiMethodName() string {
    return "alibaba.footscan.mini.image.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniImageUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniImageUploadRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlibabaFootscanMiniImageUploadRequest) GetToken() string {
    return r.token
}
// ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniImageUploadRequest) SetReqData(reqData *CheckParam) error {
    r.reqData = reqData
    r.Set("req_data", reqData)
    return nil
}

// ReqData Getter
func (r AlibabaFootscanMiniImageUploadRequest) GetReqData() *CheckParam {
    return r.reqData
}
