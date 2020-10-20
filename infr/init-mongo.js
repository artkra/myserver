db.createUser(
    {
        user: "herodotus",
        pwd: "herodotus",
        roles: [
            {
                role: "readWrite",
                db: "herodotus"
            }
        ]
    }
)