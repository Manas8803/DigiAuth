CREATE TABLE Wallet (
    id SERIAL PRIMARY KEY,
    did VARCHAR(255),
    email VARCHAR(255),
    FOREIGN KEY (email) REFERENCES users(email)
);


CREATE TABLE Certificate (
    id SERIAL PRIMARY KEY,
    walletId INTEGER,
    credentialId VARCHAR(255),
    issueDate DATE,
    FOREIGN KEY (walletId) REFERENCES Wallet(id)
);
