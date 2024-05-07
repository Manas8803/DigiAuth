CREATE TABLE Wallet (
    id SERIAL PRIMARY KEY NOT NULL,
    did VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    FOREIGN KEY (email) REFERENCES users(email)
);

CREATE TABLE Certificate (
    id SERIAL PRIMARY KEY,
    walletId INTEGER,
    credentialId VARCHAR(255),
    issueDate DATE,
    FOREIGN KEY (walletId) REFERENCES Wallet(id)
);
