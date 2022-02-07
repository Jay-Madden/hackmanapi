package data

const createUserTable = `CREATE TABLE IF NOT EXISTS "Users"
	(
		"Id" 				SERIAL NOT NULL,
		"UserName"			VARCHAR(100),
		"ApiKey"			VARCHAR(10),
		CONSTRAINT "Pk_Users" PRIMARY KEY ("Id"),
		CONSTRAINT "ApiKey" UNIQUE ("ApiKey")
	);`

const createRequestTable = `CREATE TABLE IF NOT EXISTS "Requests"
	(
		"Id" 						SERIAL NOT NULL,
		"UserId"					INTEGER NOT NULL,
		"ReturnedWord"				VARCHAR(50),
		"Length"					VARCHAR(15),
		"Time"						TIMESTAMP WITHOUT TIME ZONE,
		CONSTRAINT "Pk_Request"  	PRIMARY KEY ("Id"),
		CONSTRAINT "Fk_User" FOREIGN KEY ("UserId")
        	REFERENCES public."Users" ("Id") MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE RESTRICT
	);`
