const App = {
	data() {
		return {
			tab: "purge_candidates",
			state: null,
		};
	},
	computed: {
		notFollowers() {
			return this.state
				? this.state.effect.not_followers.sort((a, b) => a.screen_name.localeCompare(b.screen_name))
				: [];
		},
		notFriends() {
			return this.state ? this.state.effect.not_friends.sort((a, b) => a.screen_name.localeCompare(b.screen_name)) : [];
		},
		purges() {
			return this.state
				? this.state.effect.purge_candidates.sort((a, b) => a.screen_name.localeCompare(b.screen_name))
				: [];
		},
		ignores() {
			return this.state ? this.state.ignore.sort((a, b) => a.localeCompare(b)) : [];
		},
		plans() {
			return this.purges.filter(ent => !this.ignores.some(ign => ign == ent.screen_name));
		},
	},
	async mounted() {
		const resp = await fetch("./state.json");
		this.state = await resp.json();
	},
};

Vue.createApp(App).mount(document.querySelector("main"));
