const client = require('./config');

const usersIndex = 'users';
const friendshipsIndex = 'friendships';

async function updateFriendship(userName) {
    try {
        const { hits } = await client.search({
            index: friendshipsIndex,
            body: {
                query: {
                    bool: {
                        should: [
                            { match: { user: userName } },
                            { match: { friends: userName } },
                        ],
                    },
                },
            },
        });

        const results = hits?.hits || [];
        if (results.length === 0) {
            console.log(`No friendships found for user "${userName}".`);
            return;
        }

        for (const result of results) {
            const docId = result._id;
            const friendship = result._source;

            if (friendship.user === userName) {
                await client.delete({
                    index: friendshipsIndex,
                    id: docId,
                    refresh: true,
                });
                console.log(`Friendship owned by "${userName}" deleted.`);
            } else {
                const updatedFriends = friendship.friends.filter((friend) => friend !== userName);
                await client.update({
                    index: friendshipsIndex,
                    id: docId,
                    doc: { friends: updatedFriends },
                    refresh: true,
                });
                console.log(`Friendship updated for document ID ${docId}, removed "${userName}" from friends.`);
            }
        }
    } catch (error) {
        console.error(`Error updating friendships for "${userName}":`, error.meta?.body || error);
    }
}

async function deleteUserByName(name) {
    try {
        const { hits } = await client.search({
            index: usersIndex,
            body: {
                query: {
                    match: { name },
                },
            },
        });

        const results = hits?.hits || [];
        if (results.length === 0) {
            console.log(`User with name "${name}" not found.`);
            return;
        }

        const docId = results[0]._id;

        const response = await client.delete({
            index: usersIndex,
            id: docId,
            refresh: true,
        });

        console.log(`User with name "${name}" deleted successfully:`, response);
    } catch (error) {
        console.error(`Error deleting user with Name ${name}:`, error.meta?.body || error);
    }
}

async function deleteUserAndFriendships(name) {
    await updateFriendship(name);
    await deleteUserByName(name);
}

(async () => {
    await deleteUserAndFriendships('Frank');
})();
