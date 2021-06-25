package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
奖池奖品查询列表 APIRequest
alibaba.benefit.query

功能：奖池奖品查询列表
业务逻辑：程序中通过奖池编号ename,业务方身份appName来查询奖池提供的奖品返回给
小程。
安全保障：为保证数据不会越权，需要卖家授，并且验证系统参数appKey。只有通过授权的，并且
appkey验证通过的，才会查数据 并透出，否则直接失败。
因为appkey是系统参数，并且程序内部可以验证appkey和业务身份appName的关系
是否一致，所以可以保证参数appName的合法性，没有越权。
*/
type AlibabaBenefitQueryRequest struct {
    model.Params

    // 奖池编号
    ename   string 

    // 商家来源身份标识（"promotion-"+appId）
    appName   string 

    // 表示奖池类型（发奖奖池传1，抽奖传0）
    awardType   string 

}

func NewAlibabaBenefitQueryRequest() *AlibabaBenefitQueryRequest{
    return &AlibabaBenefitQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaBenefitQueryRequest) GetApiMethodName() string {
    return "alibaba.benefit.query"
}

func (r AlibabaBenefitQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaBenefitQueryRequest) SetEname(ename string) error {
    r.ename = ename
    r.Set("ename", ename)
    return nil
}

func (r AlibabaBenefitQueryRequest) GetEname() string {
    return r.ename
}

func (r *AlibabaBenefitQueryRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

func (r AlibabaBenefitQueryRequest) GetAppName() string {
    return r.appName
}

func (r *AlibabaBenefitQueryRequest) SetAwardType(awardType string) error {
    r.awardType = awardType
    r.Set("award_type", awardType)
    return nil
}

func (r AlibabaBenefitQueryRequest) GetAwardType() string {
    return r.awardType
}

