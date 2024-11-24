const users = require('../../shared/users.json');
const friendships = require('../../shared/friendships.json');
const client = require("../config/elastic");

async function init() {
    try {
        for (const index of ['users', 'friendships']) {
            if (await client.indices.exists({index})) {
                console.log(`Index "${index}" exists. Deleting...`);
                await client.indices.delete({index});
            }
        }

        console.log('Creating "users" index...');
        await client.indices.create({
            index: 'users',
            body: {
                mappings: {
                    properties: {
                        name: {type: 'text', fields: {keyword: {type: 'keyword'}}},
                        age: {type: 'integer'},
                        hobbies: {type: 'text', fields: {keyword: {type: 'keyword'}}},
                    },
                },
            },
        });

        console.log('Creating "friendships" index...');
        await client.indices.create({
            index: 'friendships',
            body: {
                mappings: {
                    properties: {
                        user: {type: 'keyword'},
                        friends: {type: 'keyword'},
                    },
                },
            },
        });

        console.log('Indexing user data...');
        for (const user of users) {
            await client.index({
                index: 'users',
                body: user,
            });
        }

        console.log('Indexing friendship data...');
        for (const friendship of friendships) {
            await client.index({
                index: 'friendships',
                body: friendship,
            });
        }

        await client.indices.refresh({index: 'users'});
        await client.indices.refresh({index: 'friendships'});

        console.log('Initialization complete.');
    } catch (error) {
        console.error('Error initializing Elasticsearch:', error.meta?.body || error);
    }
}

init().then(() => {
    console.log('Script execution finished.');
});
