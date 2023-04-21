package util

type REQ struct {
	query string // 正文
	from  string // 源
	to    string // 目的
	appid string
	salt  string // 随机数
	sign  string //
	key   string // 密钥
}

func (r *REQ) SetKey(k string) {
	r.key = k
}
func (r *REQ) SetQuery(q string) {
	r.query = q
}

func (r *REQ) SetFrom(f string) {
	r.from = f
}
func (r *REQ) SetTo(t string) {
	r.to = t
}
func (r *REQ) SetAppid(id string) {
	r.appid = id
}
func (r *REQ) SetSalt(random string) {
	r.salt = random
}
func (r *REQ) SetSign(s string) {
	r.sign = s
}
