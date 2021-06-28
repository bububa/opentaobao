package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖接口 APIResponse
alibaba.benefit.draw

功能：抽奖功能，供小程序抽奖调用
业务逻辑：程序中通过奖池编号ename,业务方身份appName来查询奖池，根据授权用户(买家)来确认抽奖用户。然后程序进行抽奖流程。
小程。
安全保障：为保证数据不会越权，需要买家授，并且验证系统参数appKey。只有通过授权的，并且
appkey验证通过的，才会进入抽奖流程，否则直接失败。
因为appkey是系统参数，并且程序内部可以验证appkey和业务身份appName的关系
是否一致，所以可以保证参数appName的合法性，没有越权。
*/
type AlibabaBenefitDrawAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaBenefitDrawResponse `json:"alibaba_benefit_draw_response,omitempty"` 
    AlibabaBenefitDrawResponse
}

/* model for simplify = false
type AlibabaBenefitDrawResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaBenefitDrawResult  *AlibabaBenefitDrawResult `json:"alibaba_benefit_draw_result,omitempty"`
    } `json:"result,omitempty"`
    

    // 权益id
    
    PrizeId   string `json:"prize_id,omitempty"`
    

    // 奖品id
    
    RightId   string `json:"right_id,omitempty"`
    

    // 扩展信息
    
    ExtAttribute   string `json:"ext_attribute,omitempty"`
    

}
*/

type AlibabaBenefitDrawResponse struct {

    // 接口返回model
    Result   *AlibabaBenefitDrawResult `json:"result,omitempty"`

    // 权益id
    PrizeId   string `json:"prize_id,omitempty"`

    // 奖品id
    RightId   string `json:"right_id,omitempty"`

    // 扩展信息
    ExtAttribute   string `json:"ext_attribute,omitempty"`

}
