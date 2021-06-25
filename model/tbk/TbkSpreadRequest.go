package tbk

// TbkSpreadRequest 
type TbkSpreadRequest struct {

    // 原始url, 只支持uland.taobao.com，s.click.taobao.com， ai.taobao.com，temai.taobao.com的域名转换，否则判错
    Url   string `json:"url,omitempty"`

}
