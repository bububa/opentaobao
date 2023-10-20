package tbk

// TaobaotbkdgnewuserordersumData 结构体
type TaobaotbkdgnewuserordersumData struct {
	// resultList
	Results []TaobaotbkdgnewuserordersumData `json:"results,omitempty" xml:"results>taobaotbkdgnewuserordersum_data,omitempty"`
	// 活动ID
	Activityid string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 日期
	Bizdate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 结算CPA 奖励金额：仅支持member 维度的统计
	Alipayusercpapreamt string `json:"alipay_user_cpa_pre_amt,omitempty" xml:"alipay_user_cpa_pre_amt,omitempty"`
	// 当日激活且首购结算的CPA 金额：仅适用于八天乐，仅支持member维度的统计
	Bindbuyusercpapreamt string `json:"bind_buy_user_cpa_pre_amt,omitempty" xml:"bind_buy_user_cpa_pre_amt,omitempty"`
	// 渠道关系id
	Relationid string `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// data
	Data *TaobaotbkdgnewuserordersumData `json:"data,omitempty" xml:"data,omitempty"`
	// 页码
	Pageno int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页大小
	Pagesize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 新注册用户数
	Regusercnt int64 `json:"reg_user_cnt,omitempty" xml:"reg_user_cnt,omitempty"`
	// 新激活用户数
	Loginusercnt int64 `json:"login_user_cnt,omitempty" xml:"login_user_cnt,omitempty"`
	// 首购用户数
	Alipayusercnt int64 `json:"alipay_user_cnt,omitempty" xml:"alipay_user_cnt,omitempty"`
	// 结算有效用户数
	Rcvvalidusercnt int64 `json:"rcv_valid_user_cnt,omitempty" xml:"rcv_valid_user_cnt,omitempty"`
	// 确认收货用户数
	Rcvusercnt int64 `json:"rcv_user_cnt,omitempty" xml:"rcv_user_cnt,omitempty"`
	// 当日激活且首购的有效用户数：仅适用于八天乐，支持member，adzone维度的统计
	Bindbuyvalidusercnt int64 `json:"bind_buy_valid_user_cnt,omitempty" xml:"bind_buy_valid_user_cnt,omitempty"`
	// bindCardValidUserCnt
	Bindcardvalidusercnt int64 `json:"bind_card_valid_user_cnt,omitempty" xml:"bind_card_valid_user_cnt,omitempty"`
	// reBuyValidUserCnt
	Rebuyvalidusercnt int64 `json:"re_buy_valid_user_cnt,omitempty" xml:"re_buy_valid_user_cnt,omitempty"`
	// validNum
	Validnum int64 `json:"valid_num,omitempty" xml:"valid_num,omitempty"`
	// 是否有下一页
	Hasnext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
