package api

import (
    "github.com/ratel-online/client/model"
    "github.com/ratel-online/core/util/http"
    "github.com/ratel-online/core/util/json"
)

const API = "http://127.0.0.1:9088"

func Login(username, password string) (*model.LoginResp, error) {
    body, err := http.Post(API+"/api/v1/login", map[string]interface{}{
        "username": username,
        "password": password,
    }, nil)
    if err != nil {
        return nil, err
    }
    resp := model.LoginResp{}
    _ = json.Unmarshal(body, &resp)
    return &resp, nil
}
