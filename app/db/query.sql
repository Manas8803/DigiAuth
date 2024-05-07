
query CreateNewWallet {
    INSERT INTO Wallet (Did, Email)
    VALUES ($1, $2)
    RETURNING *;
}

query CreateNewCertificate {
    INSERT INTO Certificate (WalletId, CredentialId, IssueDate)
    VALUES ($1, $2, $3)
    RETURNING *;
}
