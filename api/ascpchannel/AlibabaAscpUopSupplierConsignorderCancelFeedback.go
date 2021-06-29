package ascpchannel

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/ascpchannel"
)

/* 
商家仓wms取消发货反馈回告服务 
alibaba.ascp.uop.supplier.consignorder.cancel.feedback

履约单纬度通知商家仓wms取消发货结果反馈回告服务
*/
func AlibabaAscpUopSupplierConsignorderCancelFeedback(clt *core.SDKClient, req *ascpchannel.AlibabaAscpUopSupplierConsignorderCancelFeedbackRequest, session string) (*ascpchannel.AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse, error) {
    var resp ascpchannel.AlibabaAscpUopSupplierConsignorderCancelFeedbackAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
