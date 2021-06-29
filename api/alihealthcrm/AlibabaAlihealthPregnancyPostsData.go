package alihealthcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alihealthcrm"
)

/* 
发回帖子信息同步 
alibaba.alihealth.pregnancy.posts.data

发回帖子信息同步
*/
func AlibabaAlihealthPregnancyPostsData(clt *core.SDKClient, req *alihealthcrm.AlibabaAlihealthPregnancyPostsDataRequest, session string) (*alihealthcrm.AlibabaAlihealthPregnancyPostsDataAPIResponse, error) {
    var resp alihealthcrm.AlibabaAlihealthPregnancyPostsDataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
