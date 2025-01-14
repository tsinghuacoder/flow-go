/*
 * Access API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type Block struct {
	Header          *BlockHeader     `json:"header"`
	Payload         *BlockPayload    `json:"payload,omitempty"`
	ExecutionResult *ExecutionResult `json:"execution_result,omitempty"`
	Expandable      *BlockExpandable `json:"_expandable,omitempty"`
	Links           *Links           `json:"_links,omitempty"`
	BlockStatus     string           `json:"block_status"`
}
