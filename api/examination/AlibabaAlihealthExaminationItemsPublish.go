package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
单项/加项包信息同步 
alibaba.alihealth.examination.items.publish

体检机构对接_单项/加项包信息发布／更新
*/
func AlibabaAlihealthExaminationItemsPublish(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationItemsPublishAPIRequest, session string) (*examination.AlibabaAlihealthExaminationItemsPublishAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationItemsPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
