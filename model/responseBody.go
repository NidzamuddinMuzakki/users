package model

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type WebResponseListAndDetail struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Info   interface{} `json:"info"`
}
type InfoDetail struct {
	NextRowId int `json:"nextRowId"`
	PrevRowId int `json:"prevRowId"`
}
type InfoList struct {
	Allrec  int `json:"allrec"`
	Sentrec int `json:"sentrec"`
}
type WebResponseError struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
}
type ReqList struct {
	Page    int    `query:"page" myvalidator:"type:stringnumber;minLength:5;maxLength:5"`
	Perpage int    `query:"perpage"`
	Filter  string `query:"filter"`
	Order   string `query:"order"`
	Header  string `query:"header"`
}

type ReqListByUsername struct {
	Username string `query:"username"`
}
type ReqListPilihan struct {
	Filter    string `query:"filter"`
	Order     string `query:"order"`
	Type      string `query:"type"`
	Condition string `query:"condition"`
	Header    string `query:"header"`
}
type ReqDetail struct {
	RowID         int    `query:"rowId"`
	BatchNo       string `query:"batch_no"`
	VoucherNo     string `query:"voucher_no"`
	NormalBalance string `query:"normal_balance"`
	GLNo          string `query:"gl_no"`
	GLOrigin      string `query:"gl_no_origin"`
	ClosureDate   string `query:"closure_date"`
	BranchTarget  string `query:"branch_id_target`
	Header        string `query:"header"`
	Filter        string `query:"filter"`
	Order         string `query:"order"`
}
type ReqDetailIntOffice struct {
	RowID         int    `query:"rowId"`
	BatchNo       string `query:"batch_no"`
	VoucherNo     string `query:"voucher_no"`
	NormalBalance string `query:"dc_code"`
	GLNo          string `query:"gl_no"`
	GLOrigin      string `query:"gl_no_origin"`
	BranchOrigin  string `query:"branch_origin"`
	BranchTarget  string `query:"branch_id_target`
	Header        string `query:"header"`
	Filter        string `query:"filter"`
	Order         string `query:"order"`
}

type ReqDbaMovement struct {
	RowID            int    `query:"rowId"`
	AccountNo        string `query:"ac_no"`
	MovementDate     string `query:"movement_date"`
	MovementDateTime string `query:"movement_datetime"`
	MovementId       string `query:"movement_id"`
	Header           string `query:"header"`
	Filter           string `query:"filter"`
	Order            string `query:"order"`
}
type ReqMDGlBalanceDetail struct {
	RowID  int    `query:"rowId"`
	GLNo   string `query:"gl_no"`
	Header string `query:"header"`
	Filter string `query:"filter"`
	Order  string `query:"order"`
}
type ReqDetailMsob struct {
	RowID      int    `query:"rowId"`
	VoucherNo  string `query:"voucher_no"`
	GLNoDebit  string `query:"gl_no_debit"`
	GLNoCredit string `query:"gl_no_credit"`
	Filter     string `query:"filter"`
	Order      string `query:"order"`
	Header     string `query:"header"`
}
type TellerMatchingAdd struct {
	UserId string `query:"user_id"`
}
