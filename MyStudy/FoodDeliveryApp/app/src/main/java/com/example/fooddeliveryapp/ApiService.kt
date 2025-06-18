package com.example.foodapp

import retrofit2.http.GET

data class SampleData(
    val id: Int,
    val title: String
)

interface ApiService {
    @GET("posts/1") // пример эндпоинта (замени на свой)
    suspend fun getSampleData(): SampleData
}
