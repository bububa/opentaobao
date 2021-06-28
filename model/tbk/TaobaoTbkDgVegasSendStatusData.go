package tbk

// TaobaoTbkDgVegasSendStatusData 
/* model for simplify = false
type TaobaoTbkDgVegasSendStatusData struct {

    // 返回结果封装对象
    
    ResultList  struct {
        TaobaoTbkDgVegasSendStatusMapData  []TaobaoTbkDgVegasSendStatusMapData `json:"taobao_tbk_dg_vegas_send_status_map_data,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

// TaobaoTbkDgVegasSendStatusData 
type TaobaoTbkDgVegasSendStatusData struct {

    // 返回结果封装对象
    ResultList   []TaobaoTbkDgVegasSendStatusMapData `json:"result_list,omitempty"`

}
