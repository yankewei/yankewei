<template>
    <div class="bg-header h-2/5 pt-10 pl-24">
        <Title/>
        <Nav/>
    </div>
</template>

<script>
    import Title from "./Title.vue"
    import Nav from "./Nav.vue"

    export default {
        name: 'Header',
        components: {
            Title,
            Nav
        }
    }
</script>

<style scoped>
    .bg-header {
        background-color: #333333;
    }
</style>