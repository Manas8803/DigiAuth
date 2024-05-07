
-- name: CreateWallet :one
INSERT INTO Wallet (Did, Email)
VALUES ($1, $2)
RETURNING *;

-- name: CreateCertificate :one
INSERT INTO Certificate (WalletId, CredentialId, IssueDate)
VALUES ($1, $2, $3)
RETURNING *;