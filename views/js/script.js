

new Vue({
    el: "#app",
    data(){
        return {
            mov_id: '',
            mov_release: '',
            img_cover: '',
            notif: '<div class="alert alert-info">' +
            'Give it a click if you like.' +
           '</div>',
        }
    },
    methods:{
        iClick : function (event){
            axios.post('', {
                mov_id: this.mov_id,
                mov_release: this.mov_release,
                img_cover: this.img_cover,



            }).then((response) => {
                this.notif = '<div class="alert alert-success">' +
                response.data.message +
               '</div>'

               this.mov_release = ''
               this.mov_id = ''
               this.img_cover = ''

               location.reload();

            }).catch((err) => {
                console.log(err)
            })
        }
    }
})






