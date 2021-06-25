package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
工单核销 APIRequest
tmall.workcard.identify

工单核销，当工单完成以后，通过调用此接口核销
可以按照多维度核销工单，
电器预约安装按照工单维度核销，必选字段workcard_id,buyer_id,identify_code，可选字段attrs，通过扩展字段attrs回传机器码，格式{sn:'机器码'}
*/
type TmallWorkcardIdentifyRequest struct {
    model.Params

    // 核销dto
    verifyRequestDTO   *VerifyRequestDto 

}

func NewTmallWorkcardIdentifyRequest() *TmallWorkcardIdentifyRequest{
    return &TmallWorkcardIdentifyRequest{
        Params: model.NewParams(),
    }
}

func (r TmallWorkcardIdentifyRequest) GetApiMethodName() string {
    return "tmall.workcard.identify"
}

func (r TmallWorkcardIdentifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallWorkcardIdentifyRequest) SetVerifyRequestDTO(verifyRequestDTO *VerifyRequestDto) error {
    r.verifyRequestDTO = verifyRequestDTO
    r.Set("verify_request_d_t_o", verifyRequestDTO)
    return nil
}

func (r TmallWorkcardIdentifyRequest) GetVerifyRequestDTO() *VerifyRequestDto {
    return r.verifyRequestDTO
}

