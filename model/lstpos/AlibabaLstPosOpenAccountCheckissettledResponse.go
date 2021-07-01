package lstpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
校验当前用户是否入驻了零售通门店接口 API返回值 
alibaba.lst.pos.open.account.checkissettled

校验当前用户是否入驻了零售通门店接口
*/
type AlibabaLstPosOpenAccountCheckissettledAPIResponse struct {
    model.CommonResponse
    AlibabaLstPosOpenAccountCheckissettledResponse
}

// 校验当前用户是否入驻了零售通门店接口 成功返回结果
type AlibabaLstPosOpenAccountCheckissettledResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_pos_open_account_checkissettled_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果对象
    Result   *AlibabaLstPosOpenAccountCheckissettledResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
