const client = require('./config');

const usersIndex = 'users';
const friendshipsIndex = 'friendships';

async function createNewUser(data) {
    try {
        const response = await client.index({
            index: usersIndex,
            document: data,
            refresh: true,
        });

        console.log(`New user created:`, response);
    } catch (error) {
        console.error('Error creating new user:', error.meta?.body || error);
    }
}

async function createNewFriendship(data) {
    try {
        const response = await client.index({
            index: friendshipsIndex,
            document: data,
            refresh: true,
        });

        console.log(`New friendship created:`, response);
    } catch (error) {
        console.error('Error creating new friendship:', error.meta?.body || error);
    }
}

(async () => {
    await createNewUser({
        name: 'Frank',
        age: 29,
        hobbies: ['gaming', 'traveling'],
    });
    await createNewFriendship({
        user: 'Frank',
        friends: ['Eve', 'Dave'],
    });
})();
