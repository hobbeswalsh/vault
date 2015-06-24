package http

import (
	"log"
	"net/http"

	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/vault"
)

func handleSysSsh(core *vault.Core) http.Handler {
	log.Printf("Vishal: http.sys_ssh.handleSysSsh!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			respondError(w, http.StatusMethodNotAllowed, nil)
			return
		}
		log.Printf("Vishal: http.sys_ssh.handleSysSsh: requesting\n")
		var req SshRequest
		if err := parseRequest(r, &req); err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		resp, ok := request(core, w, r, requestAuth(r, &logical.Request{
			Operation: logical.WriteOperation,
			Path:      "ssh/creds/web",
			Data: map[string]interface{}{
				"username": req.Username,
				"address":  req.Address,
			},
		}))
		if !ok {
			return
		}
		respondOk(w, resp.Data)
	})
}

type SshRequest struct {
	Username string `json: "username"`
	Address  string `json: "address"`
}