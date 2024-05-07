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

CREATE TABLE users (
    id         bigserial PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password   text NOT NULL,
    isverified BOOLEAN NOT NULL DEFAULT false,
    role       text CHECK (role IN ('Issuer', 'User', 'Verifier')) NOT NULL,
    otp        text NOT NULL
    CONSTRAINT valid_email CHECK (email ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$')
);
