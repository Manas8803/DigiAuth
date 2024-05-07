
-- name: CreateWallet :one
INSERT INTO Wallet (Did, Email)
VALUES ($1, $2)
RETURNING *;

-- name: CreateCertificate :one
INSERT INTO Certificate (WalletId, CredentialId, IssueDate)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;