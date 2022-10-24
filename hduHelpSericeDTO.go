package hduHelpServiceSDK

type HduhelpBaseResponse struct {
	Cache bool   `json:"cache"`
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

type GetStudentInfoResponse struct {
	HduhelpBaseResponse
	Data struct {
		ClassId     string `json:"classId"`
		MajorId     string `json:"majorId"`
		MajorName   string `json:"majorName"`
		StaffId     string `json:"staffId"`
		StaffName   string `json:"staffName"`
		TeacherId   string `json:"teacherId"`
		TeacherName string `json:"teacherName"`
		UnitId      string `json:"unitId"`
		UnitName    string `json:"unitName"`
	} `json:"data"`
}

type GetTokenResponse struct {
	HduhelpBaseResponse
	Data struct {
		AccessToken        string `json:"access_token,omitempty"`
		AccessTokenExpire  int64  `json:"access_token_expire,omitempty"`
		RefreshToken       string `json:"refresh_token,omitempty"`
		RefreshTokenExpire int64  `json:"refresh_token_expire,omitempty"`
		StaffId            string `json:"staff_id,omitempty"`
	} `json:"data"`
}
