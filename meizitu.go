package main
// 这是个写给孔少爷的黄图爬虫
import(
   "fmt"
   "io/ioutil"
   "net/http"
   "strings"
   "os"
   "math/rand"
   "time"
   "strconv"
)

var file_path = "images/"
var dir_name = "images"

func Get_user_agent()(user_agent_list map[int]string){
    var user_agent map[int]string
    user_agent = make(map[int]string)
    user_agent[0] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3"
    user_agent[1] = "Mozilla/5.0 (X11; CrOS i686 2268.111.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11"
    user_agent[2] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3"
    user_agent[3] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3"
    user_agent[4] = "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24"
    user_agent[5] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3"
    user_agent[6] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3"
    user_agent[7] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3"
    user_agent[8] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"
    user_agent[9] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36(KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"
    user_agent[10] = "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/19.77.34.5 Safari/537.1"
    user_agent[11] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3"
    user_agent[12] = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3"
    user_agent[13] = "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3"
    user_agent[14] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/535.24 (KHTML, like Gecko) Chrome/19.0.1055.1 Safari/535.24"
    user_agent[15] = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1"
    user_agent[16] = "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6"
    user_agent[17] = "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.36 Safari/536.5"
    user_agent[18] = "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3"
    user_agent[19] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1"
    user_agent[20] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6"
    user_agent[21] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.9 Safari/536.5"
    user_agent[22] = "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3"
    user_agent[23] = "Mozilla/5.0 (Windows NT 6.2) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3"
    user_agent_list = user_agent
    return
}
// 构建浏览器头文件

func choose_agent() ( user_agent string){
    var agent_list map[int]string = Get_user_agent()
    var random_index int = rand.Intn(24)
    user_agent = agent_list[random_index]
    return
    // 随机选取一个浏览器头
}

func make_dir(dir_name string){
    err := os.Mkdir(dir_name, 0777)
    if err != nil {
    fmt.Println("mkdir root failed!")
    return
    // 建一个文件夹用来保存图片
   }
}

func get_url(url string) (content string){
    client := &http.Client{}
    request,err_1 := http.NewRequest("GET", url, nil)
    if err_1 != nil {
        fmt.Println(err_1)
        content = "0"
        return
    }
    // 设置一个模拟客户端，确认访问方式和网址
    var user_agent string = choose_agent()
    // 从浏览器信息列表中随机选一个
    request.Header.Add("User-Agent",user_agent)
    request.Header.Add("Referer", url)
    // 添加请求头，尤其是浏览器信息
    response, err_2 := client.Do(request)
    // 进行请求
    if err_2 != nil {
        fmt.Println(err_2)
        time.Sleep(time.Duration(600)*time.Second)
        content = "0"
        return
    }
    defer response.Body.Close()
    data, err_4 := ioutil.ReadAll(response.Body)
    if err_4 != nil {
        fmt.Println("read error")
        content = "0"
        return
    }
    content = string(data)
    // 读取页面的html信息
    return
}   

func get_comment_url(url string)(comment_list []string){
    var html string
    html = get_url(url)
    // 访问提供的网址
    html_list := strings.Split(html, "<li>")
    var length_html_list int = len(html_list)
    // 对得到的html进行切片操作，获取其中的li标签
    content_list := []string {}
    for _,content := range html_list[1:length_html_list-2]{
        url_blank := strings.Split(strings.Split(strings.Split(content, "<span>")[0]," ")[1],"\"")[1]
        content_list = append(content_list,url_blank)
        // 获取评论列表
    }
    comment_list = content_list
    // 返回评论列表
    return 
}

func get_page_range(url string)(max_page int){
    var html string
    var page_range string
    var max_page_content string
    var max_page_num int
    html = get_url(url)
    page_range = strings.Split(html,"class=\"pagination\"")[1]
    page_list := strings.Split(strings.Split(page_range,"</div>")[0],"</a>")
    max_page_content = strings.Split(page_list[len(page_list)-2],"\"")[1]
    max_page_num = len(max_page_content)
    max_page,err := strconv.Atoi(max_page_content[:max_page_num-4])
    if err != nil{
        fmt.Println("类型转换失败")
        return
    }
    return
}
// 获取总共多少页

func get_image_list(url string,c chan int){
    var html string
    html = get_url(url)
    image_list := []string {}
    if "0" != html{
        pic_html_1 := strings.Split(html,"<div class=\"content\">")
        if len(pic_html_1) >= 2{   
            pic_html_2 := strings.Split(pic_html_1[1],"<br>")
            length_pic_html := len(pic_html_2)
            pic_list := pic_html_2[:length_pic_html-1]
            // 获得网页中包含图片地址的列表
            for _,i := range pic_list{
                pic_url_content := strings.Split(i,"<br />")
                length_pic_url := len(pic_url_content)
                pic_index := strings.Split(pic_url_content[length_pic_url-1],"\"")
                pic_index_index := len(pic_index)
                pic_url := pic_index[pic_index_index-2]
                image_list = append(image_list,pic_url)
            // 对列表中的每个元素进行处理
            }
            download_pic(image_list)
        }   
    }
    c<-1
    // 多线程下载图片，同时下载若干个imagelist中的url
}

func download_pic(pic_url_list []string){
    for _,pic_url := range pic_url_list{
        file_name_list := strings.Split(pic_url,"/")
        file_name := file_name_list[len(file_name_list)-1]
        content := get_url(pic_url)
        if content != "0"{
            ioutil.WriteFile(file_path+file_name,[]byte(content),0666)
            fmt.Println("Download ==> ",pic_url,"successful")
        }
    }
    // 下载图片的方法，把图片打开后转为字节码，写入文件中
}

func run(url string,init_url string,c chan int){
    comment_list := get_comment_url(url)
    // 获取评论列表
    for _,comment := range comment_list{
        comment_url := init_url + comment
        get_image_list(comment_url,c)
    // 提取每个评论列表中的图片网址，组成图片网址列表，开始下载
    }
    c <- 1
    // 主要的运行流程
}

func main(){
    var init_url string
    make_dir(dir_name)
    fmt.Printf("input the index url:(defult:\"https://www.ttt443.com\"):")
    fmt.Scanln(&init_url)
    if init_url == ""{
        init_url = "https://www.ttt443.com"
    }
    // 获取输入的网址，如果没有输入，则使用默认的网址
    page_index := []string{"1","2","3","4","6","7","8"}
    var url_string string
    url_list := []string {}
    var url string
    for _,page := range page_index{
        url = "https://www.ttt443.com/htm/piclist"+page+"/"
        max_page := get_page_range(url)
        // 获取每一个对应的图片网址的索引
        for i := 1;i<= max_page;i++{
            url_string = url+strconv.Itoa(i)+".htm"
            url_list = append(url_list,url_string)
        }
    }
    // 获取全部单页的列表
    ch:=make(chan int, len(url_list)-1)
    for _,url := range url_list{
        run(url,init_url,ch)  
    }
    // 通过并发进行运行，同时开始爬取图片
    var sum int
    sum=0
    forEnd:
    for{
        select{
        case <- ch:
            sum+=1
            if sum == len(url_list)-1{
                break forEnd
            }
        }
    }   
}
























