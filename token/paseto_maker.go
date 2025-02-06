package token

import "github.com/o1egl/paseto"

type PasetoMaker struct {
	paseto       *paseto.V2
	symetric_key []byte
}

func NewPasetoMaker(symetric_key string) (*PasetoMaker, error) {
	paseto := paseto.NewV2()

	return &PasetoMaker{
		paseto:       paseto,
		symetric_key: []byte(symetric_key),
	}, nil
}

func (p *PasetoMaker) CreateToken(payload *Payload) (string, error) {
	return p.paseto.Encrypt(p.symetric_key, payload, nil)
}

func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := p.paseto.Decrypt(token, p.symetric_key, payload, nil)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
