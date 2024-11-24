const client = require("./config/elastic");


const usersIndex = 'users';
const friendshipsIndex = 'friendships';

async function testQuery() {
    try {
        const {hits} = await client.search({
            index: usersIndex,
            body: {
                query: {match: {name: 'Frank'}},
            },
        });

        const results = hits?.hits || [];
        const data = results[0]._source;
        console.log(data);
    } catch (error) {
        console.error("Error testing query:", error.meta?.body || error);
    }
}

(async () => {
    await testQuery();
})();
