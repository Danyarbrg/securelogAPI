package com.example.foodapp

import retrofit2.http.GET
import retrofit2.http.Query

data class MealResponse(
    val meals: List<Meal>?
)

data class Meal(
    val idMeal: String,
    val strMeal: String,
    val strCategory: String?,
    val strMealThumb: String,
    val strInstructions: String
)

interface MealApi {
    // Поиск по названию блюда
    @GET("api/json/v1/1/search.php")
    suspend fun searchMealsByName(
        @Query("s") name: String
    ): MealResponse
}
