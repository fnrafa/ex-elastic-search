const client = require("../config/elastic");
const usersIndex = 'users';

async function update(name, updatedData) {
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

        const response = await client.update({
            index: usersIndex,
            id: docId,
            doc: updatedData,
            refresh: true,
        });

        console.log(`User updated successfully:`, response);
    } catch (error) {
        console.error(`Error updating user with Name ${name}:`, error.meta?.body || error);
    }
}

(async () => {
    await update('Frank', {
        age: 30,
        hobbies: ['hiking', 'coding'],
    });
})();
