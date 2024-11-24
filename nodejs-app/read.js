const client = require('./config');

const usersIndex = 'users';
const friendshipsIndex = 'friendships';

async function getUserDetails() {
    try {
        const {hits} = await client.search({
            index: usersIndex, body: {
                query: {
                    match: {name: 'Frank'},
                },
            },
        });

        const results = hits?.hits || [];
        if (results.length === 0) {
            console.log(`User "Frank" not found.`);
            return;
        }

        const user = results[0]._source;
        console.log(`Details of user "Alice":`, user);
    } catch (error) {
        console.error('Error fetching user details:', error.meta?.body || error);
    }
}

async function searchUsers() {
    try {
        const {hits} = await client.search({
            index: usersIndex, body: {
                query: {
                    bool: {
                        must: [{range: {age: {gte: 25, lte: 35}}}, {match: {hobbies: 'hiking'}},],
                    },
                },
            },
        });

        const results = hits?.hits || [];
        if (results.length === 0) {
            console.log('No users match the given conditions.');
            return;
        }

        const users = results.map((hit) => hit._source);
        console.log(`Users aged between 25 and 35 with hobby "hiking":`, users);
    } catch (error) {
        console.error('Error searching users:', error.meta?.body || error);
    }
}

async function searchFriendsWithHobby() {
    try {
        const {hits: friendshipsHits} = await client.search({
            index: friendshipsIndex, body: {
                query: {
                    match: {user: 'Alice'},
                },
            },
        });

        const friendshipsResults = friendshipsHits?.hits || [];
        if (friendshipsResults.length === 0) {
            console.log(`User "Alice" not found in friendships.`);
            return;
        }

        const friends = friendshipsResults[0]._source.friends || [];
        if (friends.length === 0) {
            console.log(`User "Alice" has no friends.`);
            return;
        }

        const {hits: usersHits} = await client.search({
            index: usersIndex, body: {
                query: {
                    bool: {
                        must: [{
                            bool: {
                                should: friends.map((friend) => ({
                                    match: {name: friend},
                                })),
                            },
                        },],
                        should: [{match: {hobbies: "hiking"}}, {match: {hobbies: "gaming"}},],
                        minimum_should_match: 1,
                    },
                },
            },
        });

        const usersResults = usersHits?.hits || [];
        const friendsWithHobby = usersResults.map((hit) => hit._source.name);

        if (friendsWithHobby.length === 0) {
            console.log(`User "Alice" has no friends with preferred hobby.`);
            return;
        }

        console.log(`Friends of Alice with the hobby "hiking":`, friendsWithHobby);
    } catch (error) {
        console.error('Error searching friends with hobby:', error.meta?.body || error);
    }
}

async function searchNonFriendsWithHobby() {
    try {
        const {hits: friendshipsHits} = await client.search({
            index: friendshipsIndex, body: {
                query: {
                    match: {user: 'Alice'},
                },
            },
        });

        const friendshipsResults = friendshipsHits?.hits || [];
        if (friendshipsResults.length === 0) {
            console.log(`User "Alice" not found in friendships.`);
            return;
        }

        const friends = friendshipsResults[0]._source.friends || [];

        const {hits: usersHits} = await client.search({
            index: usersIndex, body: {
                query: {
                    bool: {
                        should: [{match: {hobbies: 'hiking'}}, {match: {hobbies: 'gaming'}},],
                        minimum_should_match: 1,
                        must_not: [{
                            bool: {
                                should: friends.map((friend) => ({
                                    match: {name: friend},
                                })),
                            },
                        }, {match: {name: 'Alice'}},],
                    },
                },
            },
        });

        const usersResults = usersHits?.hits || [];
        const nonFriendsWithHobby = usersResults.map((hit) => hit._source.name);

        console.log(`Non-friends of Alice with the hobby "hiking":`, nonFriendsWithHobby);
    } catch (error) {
        console.error('Error searching non-friends with hobby:', error.meta?.body || error);
    }
}

(async () => {
    await getUserDetails();
    await searchUsers();
    await searchFriendsWithHobby();
    await searchNonFriendsWithHobby();
})();
