// Code generated by rpcgen; DO NOT EDIT.
//
// Source: rpc/examples.go
// Template: cmd/rpcgen/examples.go.tmpl

package main

import (
	"encoding/json"

	"github.com/october93/engine/rpc"
)

var examples map[string]string = make(map[string]string)

func loadExamples() error {
	var b []byte
	var err error
	b, err = json.Marshal(rpc.AuthParamsExample1)
	if err != nil {
		return err
	}
	examples["AuthParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.AuthResponseExample1)
	if err != nil {
		return err
	}
	examples["AuthResponseExample1"] = string(b)
	b, err = json.Marshal(rpc.AuthParamsExample2)
	if err != nil {
		return err
	}
	examples["AuthParamsExample2"] = string(b)
	b, err = json.Marshal(rpc.AuthResponseExample2)
	if err != nil {
		return err
	}
	examples["AuthResponseExample2"] = string(b)
	b, err = json.Marshal(rpc.ResetPasswordParamsExample1)
	if err != nil {
		return err
	}
	examples["ResetPasswordParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.ResetPasswordResponseExample1)
	if err != nil {
		return err
	}
	examples["ResetPasswordResponseExample1"] = string(b)
	b, err = json.Marshal(rpc.LogoutParamsExample1)
	if err != nil {
		return err
	}
	examples["LogoutParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.LogoutResponseExample1)
	if err != nil {
		return err
	}
	examples["LogoutResponseExample1"] = string(b)
	b, err = json.Marshal(rpc.DeleteCardParamsExample1)
	if err != nil {
		return err
	}
	examples["DeleteCardParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.DeleteCardResponseExample1)
	if err != nil {
		return err
	}
	examples["DeleteCardResponseExample1"] = string(b)
	b, err = json.Marshal(rpc.FollowUserParamsExample1)
	if err != nil {
		return err
	}
	examples["FollowUserParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.FollowUserResponseExample1)
	if err != nil {
		return err
	}
	examples["FollowUserResponseExample1"] = string(b)
	b, err = json.Marshal(rpc.UnfollowUserParamsExample1)
	if err != nil {
		return err
	}
	examples["UnfollowUserParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.UnfollowUserResponseExample1)
	if err != nil {
		return err
	}
	examples["UnfollowUserResponseExample1"] = string(b)
	b, err = json.Marshal(rpc.GetFollowingUsersParamsExample1)
	if err != nil {
		return err
	}
	examples["GetFollowingUsersParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.GetFollowingUsersResponseExample1)
	if err != nil {
		return err
	}
	examples["GetFollowingUsersResponseExample1"] = string(b)
	b, err = json.Marshal(rpc.AddToWaitlistParamsExample1)
	if err != nil {
		return err
	}
	examples["AddToWaitlistParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.AddToWaitlistResponseExample1)
	if err != nil {
		return err
	}
	examples["AddToWaitlistResponseExample1"] = string(b)
	b, err = json.Marshal(rpc.ValidateUsernameParamsExample1)
	if err != nil {
		return err
	}
	examples["ValidateUsernameParamsExample1"] = string(b)
	b, err = json.Marshal(rpc.ValidateUsernameResponseExample1)
	if err != nil {
		return err
	}
	examples["ValidateUsernameResponseExample1"] = string(b)
	return nil
}
