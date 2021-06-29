package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
奖池奖品查询列表 API请求
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
    _ename   string
    // 商家来源身份标识（"promotion-"+appId）
    _appName   string
    // 表示奖池类型（发奖奖池传1，抽奖传0）
    _awardType   string
}

// 初始化AlibabaBenefitQueryRequest对象
func NewAlibabaBenefitQueryRequest() *AlibabaBenefitQueryRequest{
    return &AlibabaBenefitQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaBenefitQueryRequest) GetApiMethodName() string {
    return "alibaba.benefit.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaBenefitQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ename Setter
// 奖池编号
func (r *AlibabaBenefitQueryRequest) SetEname(_ename string) error {
    r._ename = _ename
    r.Set("ename", _ename)
    return nil
}

// Ename Getter
func (r AlibabaBenefitQueryRequest) GetEname() string {
    return r._ename
}
// AppName Setter
// 商家来源身份标识（"promotion-"+appId）
func (r *AlibabaBenefitQueryRequest) SetAppName(_appName string) error {
    r._appName = _appName
    r.Set("app_name", _appName)
    return nil
}

// AppName Getter
func (r AlibabaBenefitQueryRequest) GetAppName() string {
    return r._appName
}
// AwardType Setter
// 表示奖池类型（发奖奖池传1，抽奖传0）
func (r *AlibabaBenefitQueryRequest) SetAwardType(_awardType string) error {
    r._awardType = _awardType
    r.Set("award_type", _awardType)
    return nil
}

// AwardType Getter
func (r AlibabaBenefitQueryRequest) GetAwardType() string {
    return r._awardType
}
