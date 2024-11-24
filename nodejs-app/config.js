const { Client } = require('@elastic/elasticsearch');

const client = new Client({
    node: 'http://localhost:9200',
    auth: {
        username: 'elastic',
        password: '04tlHo4AJhYFhsNjC5e+',
    },
});

module.exports = client;
