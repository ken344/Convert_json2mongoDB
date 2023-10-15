var user = {
  user: "mongo",
  pwd: "mongo",
  roles: [
    {
      role: "dbOwner",
      db: "todofuken-db"
    }
  ]
};

db.createUser(user);
db.createCollection('todofuken-data');
