<template>
    <div class="flex mt-10">
        <div v-for="nav in nav_list" :key="nav.link">
            <a :href="'./' + nav.link">
                <div class="py-1 px-3 text-white rounded-md hover:bg-gray-100 hover:text-black">{{ nav.name }}</div>
            </a>
        </div>
    </div>
</template>

<style scoped>

</style>

<script>
export default {
    name: "Nav",
    data: function () {
        return {
            nav_list: [
                {
                    name: '主页',
                    link: 'home',
                },
                {
                    name: '编程',
                    link: 'tech',
                },
                {
                    name: '随笔',
                    link: 'feeling'
                }
            ]
        }
    }
}
</script>